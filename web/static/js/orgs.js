"use strict";

var orgsDropdownButton

var orgsDropdownMenu;

$(document).ready(() => {
    orgsDropdownButton = document.getElementById("orgs-dropdown-button");

    orgsDropdownMenu = document.getElementById("orgs-dropdown-menu");

    $("#orgs-dropdown-button").click(() => {
        orgs();
        $(orgsDropdownMenu).toggleClass("show");
    })
})

function orgs(){
    fetch("/orgs", {method: "GET"})
        .then(response => response.json())
        .then(data => renderOrgs(data))
}

function renderOrgs(orgsData){
    let liOrgsList = orgsData.map((org) => {
        console.log(org["login"])
        let listItem = document.createElement("div");
        listItem.id = "org-" + org["login"];
        
        let pElem = document.createElement("p");
        pElem.className = "dropdown-item list-group-item list-group-item-secondary";
        pElem.innerHTML = org["login"];

        listItem.appendChild(pElem);

        return listItem
    })

    let divOrgsList = document.getElementById("orgs-dropdown-menu");

    while (divOrgsList.firstChild){
        divOrgsList.removeChild(divOrgsList.firstChild);
    }

    liOrgsList.forEach(orgItem => divOrgsList.appendChild(orgItem));
}

