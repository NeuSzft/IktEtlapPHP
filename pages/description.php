<!DOCTYPE html>
<html lang="hu">
    
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <?php
    include_once('./data/menu-cards.php');

    $query;
    parse_str(parse_url($_SERVER['REQUEST_URI'], PHP_URL_QUERY), $query);
    $id = $query['dish'];
    echo "<title>$id</title>\n"
    ?>
    <link rel="stylesheet" href="../style/style.css">
    <link rel="stylesheet" href="../style/description.css">

</head>

<body>
    <nav id="desktop-nav" class="navbar">
        <a href="/?day=hetfo">Hétfő</a>
        <a href="/?day=kedd">Kedd</a>
        <a href="/?day=szerda">Szerda</a>
        <a href="/?day=csutortok">Csütörtök</a>
        <a href="/?day=pentek">Péntek</a>
        <a href="/?day=szombat">Szombat</a>
        <a href="/?day=vasarnap">Vasárnap</a>
    </nav>

    <nav id="mobile-nav" class="navbar">
        <div id="days">
            <a href="/?day=hetfo">Hétfő</a>
            <a href="/?day=kedd">Kedd</a>
            <a href="/?day=szerda">Szerda</a>
            <a href="/?day=csutortok">Csütörtök</a>
            <a href="/?day=pentek">Péntek</a>
            <a href="/?day=szombat">Szombat</a>
            <a href="/?day=vasarnap">Vasárnap</a>
        </div>
        <a href="javascript:void(0);" class="icon" onclick="ToggleHamburgerMenu()">
            <i class="fa fa-bars"></i>
        </a>
    </nav>
    <?php
        echo "<div>\n";
        echo "<h1>{$cards[$id]['name']}</h1>\n";
        echo "<img src=\"img/{$cards[$id]['img']}\" alt=\"{$cards[$id]['name']} kép\">\n";
        foreach ($cards[$id]['desc'] as $desc) {
            echo "<p>{$desc}</p>\n";
        }
        echo "</div>";
        ?>
    <script src="../script/script.js"></script>
</body>
</html>