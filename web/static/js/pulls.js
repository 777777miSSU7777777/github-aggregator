"use strict";

$(document).ready(() => {
    AllPulls();
    $(document).on("click", "#all-tab", () => {
        AllPulls();
    });
    $(document).on("click", "#assigned-tab", () => {
        AssignedPulls();
    });
    $(document).on("click", "#review-requested-tab", () => {
        ReviewRequestedPulls();
    });

    $("#all-tab").click();
});

const AllPulls = () => {
    fetch("/api/pulls?filter=all&selected_orgs=" + JSON.stringify( getOrgsChoiceArray() ),  { method: "GET"})
        .then(response => response.json())
        .then(data => renderAllPulls(data));
}

const renderAllPulls = (data) => {
    let pullsData = data["filtered_pulls"];

    let allTabContent = document.getElementById("all");

    if (allTabContent.firstElementChild) {
        allTabContent.removeChild(allTabContent.firstElementChild);
    }

    if (pullsData.length) {
        let pullsTab = createPullsTab(pullsData);
    
        allTabContent.appendChild(pullsTab);
    } else {
        let pItem = document.createElement("p");
        pItem.innerHTML = "There no pulls assigned or review requested to you."

        allTabContent.appendChild(pItem);
    }
}

const AssignedPulls = () => {
    fetch("/api/pulls?filter=assignee&selected_orgs=" + JSON.stringify( getOrgsChoiceArray() ),  { method: "GET"})
    .then(response => response.json())
    .then(data => renderAssignedPulls(data));
}

const renderAssignedPulls = (data) => {
    let pullsData = data["filtered_pulls"];

    let assignedTabContent = document.getElementById("assigned");

    if (assignedTabContent.firstElementChild) {
        assignedTabContent.removeChild(assignedTabContent.firstElementChild);
    }

    if (pullsData.length) {
        let pullsTab = createPullsTab(pullsData);
    
        assignedTabContent.appendChild(pullsTab);
    } else {
        let pItem = document.createElement("p");
        pItem.innerHTML = "There no pulls assigned to you."

        assignedTabContent.appendChild(pItem);
    }
}

const ReviewRequestedPulls = () => {
    fetch("/api/pulls?filter=reviewer&selected_orgs=" + JSON.stringify( getOrgsChoiceArray() ),  { method: "GET"})
        .then(response => response.json())
        .then(data => renderReviewRequestedPulls(data));
}

const renderReviewRequestedPulls = (data) => {
    let pullsData = data["filtered_pulls"];

    let reviewRequestTabContent = document.getElementById("review-requested");

    if (reviewRequestTabContent.firstElementChild) {
        reviewRequestTabContent.removeChild(reviewRequestTabContent.firstElementChild);
    }

    if (pullsData.length) {
        let pullsTab = createPullsTab(pullsData);
    
        reviewRequestTabContent.appendChild(pullsTab);
    } else {
        let pItem = document.createElement("p");
        pItem.innerHTML = "There no pulls review requested to you."

        reviewRequestTabContent.appendChild(pItem);
    }
}

const getOrgsChoiceArray = () => {
    let orgsChoice = JSON.parse( localStorage.getItem("orgs_choice") );
    let orgsChoiceArray = []

    for (let org in orgsChoice) {
        orgsChoiceArray.push(org);
    }

    return orgsChoiceArray;
}

const createPullsTab = (data) => {
    let pullsTab = document.createElement("table");
    pullsTab.className = "table";

    let tableHead = document.createElement("thead");
    pullsTab.appendChild(tableHead);
    let headRow = document.createElement("tr");
    tableHead.appendChild(headRow);

    let repoHead = document.createElement("th")
    repoHead.innerHTML = "Repository";
    headRow.appendChild(repoHead);

    let prHead = document.createElement("th");
    prHead.innerHTML = "Pull Request";
    headRow.appendChild(prHead);

    let authorHead = document.createElement("th");
    authorHead.innerHTML = "Author";
    headRow.appendChild(authorHead);

    let tableBody = document.createElement("tbody");
    pullsTab.appendChild(tableBody);

    let trList = data.map(pull => {
        let trItem = document.createElement("tr");

        let repoItem = document.createElement("td");  
        let repoLink = document.createElement("a");
        repoLink.innerHTML = pull["head"]["repo"]["full_name"];
        repoLink.href = pull["head"]["repo"]["html_url"];
        repoItem.appendChild(repoLink);
        trItem.appendChild(repoItem);

        let prItem = document.createElement("td");  
        let prLink = document.createElement("a");
        prLink.innerHTML = pull["title"];
        prLink.href = pull["html_url"];
        prItem.appendChild(prLink);
        trItem.appendChild(prItem);

        let authorItem = document.createElement("td");  
        let authorLink = document.createElement("a");
        authorLink.innerHTML = pull["user"]["login"];
        authorLink.href = pull["user"]["html_url"];
        authorItem.appendChild(authorLink);
        trItem.appendChild(authorItem);

        return trItem;
    });

    trList.forEach(pull => tableBody.appendChild(pull));

    return pullsTab;
}