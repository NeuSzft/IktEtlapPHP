<?php

$path = $_SERVER['REQUEST_URI'];

function route(string $file, string $url)
{
    if (!include('./pages/' . $file)) {
        echo "404 Not Found '$url'";
        http_response_code(404);
    }
}

switch ($path) {
    case '/desc':
        route('description.php', $path);
        break;

    default:
        route('main.php', $path);
}
