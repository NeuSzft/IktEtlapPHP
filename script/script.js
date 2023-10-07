function ToggleHamburgerMenu() {
    var links = document.getElementById("days");
    if (links.style.display === "block") {
      links.style.display = "none";
    } 
    else {
      links.style.display = "block";
    }
  }