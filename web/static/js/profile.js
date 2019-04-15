"use strict";

var profileDropdownButton;

var profileDropdownMenu;

$(document).ready(() => {
    profileDropdownButton = document.getElementById("profile-dropdown-button");

    profileDropdownMenu = document.getElementById("profile-dropdown-menu");
    
    profile();
});

function profile(){
    fetch("/profile",{method: "GET"})
        .then(response => response.json())
        .then(data => renderProfile(data));
}

function renderProfile(profileData){
    let profilePic = document.createElement("img");
    profilePic.className = "profile-pic";
    profilePic.src = profileData["avatar_url"];

    profileDropdownButton.appendChild(profilePic);

    let profileName = document.createElement("p");
    profileName.className = "profile-name";
    profileName.innerHTML = profileData["login"];

    profileDropdownButton.append(profileName);

    let seeProfileButton = document.getElementById("see-profile-button");

    seeProfileButton.addEventListener("click", () => window.open(profileData["html_url"]));

    let logoutButton = document.getElementById("logout-button");
    logoutButton.addEventListener("click", logoutAction);
}


function logoutAction(){
    let req = new XMLHttpRequest();
    req.open("POST", "/logout");
    req.send(null);
    localStorage.removeItem("orgs_choice");
    document.location.reload();
}