"use strict";

var profileDropdownButton;

var profileDropdownMenu;

$(document).ready(() => {
    profileDropdownButton = document.getElementById("profile-dropdown-button");

    profileDropdownMenu = document.getElementById("profile-dropdown-menu");
    
    profile();
})

function profile(){
    fetch("/profile",{method: "GET"})
        .then(response => response.json())
        .then(data => renderProfile(data));
}

function renderProfile(profileData){
    let profilePic = document.createElement("img");
    profilePic.className = "profile-pic"
    profilePic.src = profileData["avatar_url"];

    profileDropdownButton.appendChild(profilePic);
    profileDropdownButton.append(profileData["login"]);

    let seeProfileButton = document.getElementById("see-profile-button");

    seeProfileButton.addEventListener("click", () => window.open(profileData["html_url"]));

    let logoutButton = document.getElementById("logout-button");
    logoutButton.addEventListener("click", logoutAction);
}


function logoutAction(){
    let req = new XMLHttpRequest();
    req.open("POST", "/logout");
    req.send(null);
    document.location.reload();
}