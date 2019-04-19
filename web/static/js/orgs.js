"use strict";

var orgsDropdownButton;

var orgsDropdownMenu;

const enabledOrgClass = "list-group-item-success";

$(document).ready(() => {
    orgsDropdownButton = document.getElementById("orgs-dropdown-button");

    fetch("/api/scopes", {method: "GET"})
        .then(response => response.json())
        .then(data => orgsCheckTokenPermissions(data, orgsTokenHasPermits, orgsTokenHasntPermits));
});

const orgs = () => {
    fetch("/api/orgs", {method: "GET"})
        .then(response => response.json())
        .then(data => renderOrgs(data));
}

const renderOrgs = (orgsData) => {
    let orgsChoice = JSON.parse(localStorage.getItem("orgs_choice"));

    let liOrgsList = orgsData.map(org => {
        let listItem = document.createElement("div");
        listItem.id = "org-" + org["login"];
        
        let pElem = document.createElement("p");
        pElem.className = "dropdown-item org-item list-group-item";

        if (org["login"] in orgsChoice){
            $(pElem).addClass(enabledOrgClass);
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

const addOrg = (org) => {
    let orgsChoice = JSON.parse(localStorage.getItem("orgs_choice"));
    orgsChoice[org] = org;
    localStorage.setItem("orgs_choice", JSON.stringify(orgsChoice));
}

const delOrg = (org) =>{
    let orgsChoice = JSON.parse(localStorage.getItem("orgs_choice"));
    delete orgsChoice[org];
    localStorage.setItem("orgs_choice", JSON.stringify(orgsChoice));
}

const orgsCheckTokenPermissions = () => (scopes, enough, notEnough){
    let splitScopes = scopes["scopes"].split(",");

    splitScopes = splitScopes.map(scope => scope.trim());
    
    if ( (splitScopes.indexOf("user")  != -1) && ((splitScopes.indexOf("read:org") != -1) || (splitScopes.indexOf("admin:org") != -1 )) ){
        enough();
    } else {
        notEnough();
    }
}

const orgsTokenHasPermits = () => {
    if (localStorage.getItem("orgs_choice") == undefined){
        localStorage.setItem("orgs_choice","{}");
    }

    $(orgsDropdownButton).parent().addClass("active").attr("data-toggle","dropdown");

    orgsDropdownMenu = document.getElementById("orgs-dropdown-menu");

    $("#orgs-dropdown-button").click(() => {
        orgs();
        $(orgsDropdownMenu).toggleClass("show");
    });

    $(document).on("click", ".org-item", event => {
        const toggleOrg = $(event.target);

        let org = toggleOrg.text();

        if ( !$(toggleOrg).hasClass(enabledOrgClass) ){
            $(toggleOrg).addClass(enabledOrgClass);
            addOrg(org)
        } else if ( $(toggleOrg).hasClass(enabledOrgClass)){
            $(toggleOrg).removeClass(enabledOrgClass)
            delOrg(org);
        }
    });
}

const orgsTokenHasntPermits = () => {
    $(orgsDropdownButton).parent()
    .addClass("inactive")
    .attr("data-toggle", "popover")
    .attr("data-placement", "bottom")
    .attr("data-trigger", "hover")
    .attr("title", "Provided scopes are not enough")
    .attr("data-content", "Organizations functionality requires at least 'user' and 'read:org' scopes.");
    
    $(document).ready(() => {
        $('[data-toggle="popover"]').popover();
    })
}




