<?php

$path = $_SERVER['REQUEST_URI'];

function route(string $file, string $url)
{
    if (!include('./pages/' . $file)) {
        echo '404 Not Found ' . $path;
    }
}

switch ($path) {
    case '/desc':
        route('description.php', $path);
        break;

    default:
        route('main.php', $path);
}
