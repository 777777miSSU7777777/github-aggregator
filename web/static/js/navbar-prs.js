"use strict";


var prsTabLink;

$(document).ready(() => {
    prsTabLink = document.getElementById("prs-tab-link");

    fetch("/api/scopes", {method: "GET"})
        .then(response => response.json())
        .then(data => prsCheckTokenPermissions(data, prsTokenHasPermits, prsTokeHasntPermits))
})

const prsCheckTokenPermissions = (scopes, enough, notEnough) => {
    let splitScopes = scopes["scopes"].split(",");

    splitScopes = splitScopes.map(scope => scope.trim());

    if ( splitScopes.indexOf("read:org") != -1 || splitScopes.indexOf("admin:org") != -1){
        enough();
    } else {
        notEnough();
    }
}

const prsTokenHasPermits = () => {
    $(prsTabLink).parent("active");

    prsTabLink.href = "/pull-requests"
}

const prsTokeHasntPermits = () => {
    $(prsTabLink).parent()
    .addClass("inactive")
    .attr("data-toggle", "popover")
    .attr("data-placement", "bottom")
    .attr("data-trigger", "hover")
    .attr("title", "Provided scopes are not enough")
    .attr("data-content", "Pull requests functionality requires at least 'read:org' scope.");
    
    $(document).ready(() => {
        $('[data-toggle="popover"]').popover();
    })
}
