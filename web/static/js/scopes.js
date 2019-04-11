"use strict";

let scopesButton = document.getElementById("scopes-button");

$(document).ready(() => {
    $("#scopes-button").click(() => {
        scopes();
        $("#providedScopes").modal("show");
    });
})

function scopes(){
    fetch('/scopes')
        .then(response => response.json())
        .then(data => renderScopes(data));
}

function renderScopes(response){
    let splitScopes = response["scopes"].split(",");
    
    let liScopesList = splitScopes.map((scope) => {
        let listItem = document.createElement("li");
        listItem.className = "list-group-item list-group-item-success";

        let pItem = document.createElement("p");
        pItem.innerHTML = scope;

        listItem.appendChild(pItem);

        console.log(scope);

        return listItem;
    })

    let ulScopesList = document.getElementById("scopes-list");

    while (ulScopesList.firstChild) {
        ulScopesList.removeChild(ulScopesList.firstChild);
    }

    liScopesList.forEach(scopeItem => ulScopesList.appendChild(scopeItem));
}
