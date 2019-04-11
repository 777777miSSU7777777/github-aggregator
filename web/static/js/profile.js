"use strict";

var dropdownButton;

var dropdownMenu;

$(document).ready(() => {
    dropdownButton = document.getElementById("profileDropdownButton");

    dropdownMenu = document.getElementById("profileDropdownMenu");
    
    profile();
})

function profile(){
    fetch("/profile",{method: "GET"})
        .then(response => response.json())
        .then(data => renderProfile(data));
    console.log("123");
}

function renderProfile(profileData){
    let profilePic = document.createElement("img");
    profilePic.className = "profile-pic"
    profilePic.src = profileData["avatar_url"];

    dropdownButton.appendChild(profilePic);
    dropdownButton.append(profileData["login"]);

    let seeProfileButton = document.getElementById("see-profile-button");

    seeProfileButton.addEventListener("click", () => window.open(profileData["html_url"]));

    // let scopesButton = document.getElementById("scopes-button");
    // scopesButton.addEventListener("click", () => console.log(scopes()));

    let logoutButton = document.getElementById("logout-button");
    logoutButton.addEventListener("click", logoutAction);
}


function logoutAction(){
    let req = new XMLHttpRequest();
    req.open("POST", "/logout");
    req.send(null);
    document.location.reload();
}