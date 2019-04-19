"use strict";

var scopesButton = document.getElementById("scopes-button");

const presentScopeClass = "list-group-item-success";

const missingScopeClass = "list-group-item-danger";

const requiredOrgsScopes = ["user", "read:org"];

const requiredOrgsScopesOpt = "admin:org";


$(document).ready(() => {
    $("#scopes-button").click(() => {
        scopes();
        $("#provided-scopes").modal("show");
    });
});

const scopes = () =>{
    fetch('/api/scopes')
        .then(response => response.json())
        .then(data => renderScopes(data));
}

const renderScopes = (response) => {
    let splitScopes = response["scopes"].split(",");

    splitScopes = splitScopes.map(scope => scope.trim())
    
    let ScopesList = document.getElementById("functionality-scopes-list");

    while (ScopesList.firstChild) {
        ScopesList.removeChild(ScopesList.firstChild);
    }

    ScopesList.appendChild( getOrgsScopesList(splitScopes) );

    ScopesList.appendChild( getPRsScopesList(splitScopes) );
}

const getOrgsScopesList = (scopes) => {
    let orgsScopesDiv = document.createElement("div");
    orgsScopesDiv.className = "orgs-scopes";
    orgsScopesDiv.id = "orgs-scopes";

    let orgsScopesListHeader = document.createElement("h2");
    orgsScopesListHeader.innerHTML = "Organizations";
    orgsScopesDiv.appendChild(orgsScopesListHeader);

    let scopesList = document.createElement("ul");
    scopesList.className = "list-group";
    scopesList.id = "orgs-scopes-list";

    let liScopesList = requiredOrgsScopes.map(scope => {
        let liItem = document.createElement("li");
        $(liItem).addClass("list-group-item");

        let pItem = document.createElement("p");
        pItem.innerHTML = scope;
        liItem.appendChild(pItem);

        if (scopes.includes(scope)){
            $(liItem).addClass(presentScopeClass);
        } else {
            $(liItem).addClass(missingScopeClass);
        }

        if (scope == "read:org" && $(liItem).hasClass(missingScopeClass)){
            if (scopes.includes(requiredOrgsScopesOpt)){
                $(liItem)
                    .removeClass(missingScopeClass)
                    .addClass(presentScopeClass)
            }
        }

        return liItem;
    });

    liScopesList.forEach(scope => scopesList.appendChild(scope));
    orgsScopesDiv.append(scopesList);

    return orgsScopesDiv;
}

const getPRsScopesList = (scopes) => {
    let prsScopesDiv = document.createElement("div");
    prsScopesDiv.className = "prs-scopes";
    prsScopesDiv.id = "prs-scopes";

    let prsScopesListHeader = document.createElement("h2");
    prsScopesListHeader.innerHTML = "Pull requests";
    prsScopesDiv.appendChild(prsScopesListHeader);

    let scopesList = document.createElement("ul");
    scopesList.className = "list-group";
    scopesList.id = "prs-scopes-list";

    let liItem = document.createElement("li");
    $(liItem).addClass("list-group-item");
    let pItem = document.createElement("p");
    pItem.innerHTML = "read:org";
    liItem.appendChild(pItem);

    $(liItem).addClass( scopes.indexOf("read:org") != -1 || scopes.indexOf("admin:org") != -1 ? presentScopeClass : missingScopeClass)

    scopesList.appendChild(liItem);
    prsScopesDiv.append(scopesList);

    return prsScopesDiv;
}
