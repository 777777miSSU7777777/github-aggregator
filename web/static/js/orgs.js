"use strict";

var orgsDropdownButton

var orgsDropdownMenu;

const enabledOrgClass = "list-group-item-success";

const disabledOrgClass = "list-group-item-secondary";

$(document).ready(() => {
    orgsDropdownButton = document.getElementById("orgs-dropdown-button");

    orgsDropdownMenu = document.getElementById("orgs-dropdown-menu");

    $("#orgs-dropdown-button").click(() => {
        orgs();
        $(orgsDropdownMenu).toggleClass("show");
    });

    $(document).on("click", ".org-item", event => {
        const toggleOrg = $(event.target);

        let org = toggleOrg.text();

        if ( $(toggleOrg).hasClass(disabledOrgClass) ){
            $(toggleOrg)
                .removeClass(disabledOrgClass)
                .addClass(enabledOrgClass);
            addOrg(org)
        } else if ( $(toggleOrg).hasClass(enabledOrgClass)){
            $(toggleOrg)
                .removeClass(enabledOrgClass)
                .addClass(disabledOrgClass);
            delOrg(org);
        }
    });
});

function orgs(){
    fetch("/orgs", {method: "GET"})
        .then(response => response.json())
        .then(data => renderOrgs(data));
}

function renderOrgs(orgsData){
    let orgsChoice = JSON.parse(localStorage.getItem("orgs_choice"));

    let liOrgsList = orgsData.map(org => {
        let listItem = document.createElement("div");
        listItem.id = "org-" + org["login"];
        
        let pElem = document.createElement("p");
        pElem.className = "dropdown-item org-item list-group-item";

        if (org["login"] in orgsChoice){
            $(pElem).addClass(enabledOrgClass);
            console.log("en")
        } else {
            $(pElem).addClass(disabledOrgClass);
            console.log("di")
        }

        pElem.innerHTML = org["login"];

        listItem.appendChild(pElem);

        return listItem;
    });

    let divOrgsList = document.getElementById("orgs-dropdown-menu");

    while (divOrgsList.firstChild){
        divOrgsList.removeChild(divOrgsList.firstChild);
    }

    liOrgsList.forEach(orgItem => divOrgsList.appendChild(orgItem));
}

function addOrg(org){
    if (localStorage.getItem("orgs_choice") === undefined){
        localStorage.setItem("orgs_choice","{}");
    }

    let orgsChoice = JSON.parse(localStorage.getItem("orgs_choice"));
    orgsChoice[org] = org;
    localStorage.setItem("orgs_choice", JSON.stringify(orgsChoice));
}

function delOrg(org){
    let orgsChoice = JSON.parse(localStorage.getItem("orgs_choice"));
    delete orgsChoice[org];
    localStorage.setItem("orgs_choice", JSON.stringify(orgsChoice));
}





