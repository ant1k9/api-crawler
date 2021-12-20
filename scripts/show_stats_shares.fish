#!/usr/bin/env fish

## helpers
function show_graph
    argparse 'i/id=' 't/ticket=' 'h/height=' 'w/width=' -- $argv
    if test "$_flag_height" = ""
        set _flag_height "16"
    end
    if test "$_flag_width" = ""
        set _flag_width "200"
    end

    echo "         [$_flag_id] $_flag_ticket"
    set -l query "
        SELECT value FROM (
            SELECT
                payload->>'<CLOSE>' AS value,
                payload->>'<DATE>' AS ts
            FROM data_share_$_flag_id
            ORDER BY payload->>'<DATE>' DESC
            LIMIT 90
        ) _
        ORDER BY ts ASC"

    psql $DATABASE_URL -qXAt -c "$query" | asciigraph -w "$_flag_width" -h "$_flag_height"
    echo
end

function show_graph_by_query
    argparse 'q/query=' 'no-wait=' 'h/height=' 'w/width=' -- $argv
    for row in (psql $DATABASE_URL -qXAt -c "$_flag_query")
        set -l id (echo "$row" | awk -F'|' '{ print $1 }')
        set -l ticket (echo "$row" | awk -F'|' '{ print $2 }')
        show_graph -i "$id" -t "$ticket" -h "$_flag_height" -w "$_flag_width"
        if not test "$_flag_no_wait" = '--no-wait'
            sleep 2
        end
    end
end

function trending_query
    argparse 'o/order-by=' 'd/days=' 't/top=' -- $argv
    if not set -q "_flag_days"
        set _flag_days "30"
    end
    if test "$_flag_top" = ""
        set _flag_top "10"
    end

    echo "
    WITH share_stats AS (
        SELECT
            plugin,
            (payload->>'<CLOSE>')::DECIMAL AS value,
            ROW_NUMBER() OVER (PARTITION BY plugin ORDER BY payload->>'<DATE>' DESC) rn
        FROM data
    ),
    shares_mapping_with_plugin AS (
        SELECT
            id,
            name,
            ticket,
            'share_' || id AS plugin
        FROM shares_mapping
    )

    SELECT
        sm.id, sm.name
    FROM shares_mapping_with_plugin sm
    INNER JOIN share_stats sh1 USING (plugin)
    INNER JOIN share_stats sh2 USING (plugin)
    WHERE
        sh1.rn = 1 AND sh2.rn = $_flag_days
    ORDER BY $_flag_order_by LIMIT $_flag_top"
end

## body
argparse \
    'h/help' \
    'show=' \
    'trending' \
    'descending' \
    'd/days=' \
    'no-wait' \
    'width=' \
    'top=' \
    'height=' -- $argv

if set -q _flag_help
    echo -n 'Usage:
    ./show_stats.fish                          # help charts for all coins
    ./show_stats.fish --width 100 --height 10  # customize width and height of charts
    ./show_stats.fish --help                   # help message
    ./show_stats.fish --show Ethereum          # make Ethereum chart
    ./show_stats.fish --trending --days 7      # make charts for top 10 trending coins
    ./show_stats.fish --descending             # make charts for top 10 descending coins
'
else if set -q _flag_show
    show_graph \
        -c "$_flag_show" \
        -w "$_flag_width" \
        -h "$_flag_height"

else if set -q _flag_trending
    set -l query (trending_query -o "sh2.value / sh1.value" -d "$_flag_days" -t "$_flag_top")
    show_graph_by_query \
        -q "$query" \
        --no-wait "$_flag_no_wait" \
        -w "$_flag_width" \
        -h "$_flag_height"

else if set -q _flag_descending
    set -l query (trending_query -o "sh1.value / sh2.value" -d "$_flag_days" -t "$_flag_top")
    show_graph_by_query \
        -q "$query" \
        --no-wait "$_flag_no_wait" \
        -w "$_flag_width" \
        -h "$_flag_height"

else
    show_graph_by_query \
        -q 'SELECT id, name FROM shares_mapping ORDER BY id' \
        --no-wait "$_flag_no_wait" \
        -w "$_flag_width" \
        -h "$_flag_height"
end
