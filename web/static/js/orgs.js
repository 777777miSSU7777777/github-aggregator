"use strict";

var orgsDropdownButton

var orgsDropdownMenu;

const enabledOrgClass = "list-group-item-success";

const disabledOrgClass = "list-group-item-secondary"

$(document).ready(() => {
    orgsDropdownButton = document.getElementById("orgs-dropdown-button");

    orgsDropdownMenu = document.getElementById("orgs-dropdown-menu");

    $(document).on("click", "#orgs-dropdown-button", () => {
        orgs();
        $(orgsDropdownMenu).toggleClass("show");
    });

    $(document).on("click", ".org-item", event => {
        const toggleOrg = $(event.target);

        if ( $(toggleOrg).hasClass(disabledOrgClass) ){
            $(toggleOrg)
                .removeClass(disabledOrgClass)
                .addClass(enabledOrgClass);
        } else if ( $(toggleOrg).hasClass(enabledOrgClass)){
            $(toggleOrg)
                .removeClass(enabledOrgClass)
                .addClass(disabledOrgClass);
        }
    });
})

function orgs(){
    fetch("/orgs", {method: "GET"})
        .then(response => response.json())
        .then(data => renderOrgs(data))
}

function renderOrgs(orgsData){
    let liOrgsList = orgsData.map(org => {
        let listItem = document.createElement("div");
        listItem.id = "org-" + org["login"];
        
        let pElem = document.createElement("p");
        pElem.className = "dropdown-item org-item list-group-item list-group-item-secondary";
        pElem.innerHTML = org["login"];

        listItem.appendChild(pElem);

        return listItem
    });

    let divOrgsList = document.getElementById("orgs-dropdown-menu");

    while (divOrgsList.firstChild){
        divOrgsList.removeChild(divOrgsList.firstChild);
    }

    liOrgsList.forEach(orgItem => divOrgsList.appendChild(orgItem));
}

