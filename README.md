# HPE-NIE-CTY-Group-2

GOLang library for HPE iLO Restful Operations<br>
(Team members: Aman Narayan Singh, Swathi BJ, Aditya Mishra, Shreesh Kulkarni)

# How this idea matters to HPE: <br>
HPE is leader in server software which comprises of <b>Integrated Lights out(iLO)</b>.  To manage iLO using Restful Redfish APIs, Golang library is the first step. Library should be independently used to develop new Golang tools to manage HPE servers.

# Phase 1
<b>Objective</b>:<br> Learning about Golang, Restful, Redfish, Client/Server etc.<br>
<b>Description:</b><br> 
<b>The learning should include:<b/><br>
•	What is Restful Architecture<br>
•	Basics of Golang.<br>
•	Basics of DMTF Redfish.<br>
•	Identify Tools needed like VS code/GoLand etc.<br>
•	Setting up development environment like ubuntu VM etc.<br>
•	Identify dependencies and design intended library.<br>

<h2>Deliverable:</h2><br> Design document of Golang library for Rest/Redfish.<br>

So hereby I am attaching the ppt for the project understanding: <a href="https://docs.google.com/presentation/d/1txHXZgsGbmgA1qwHUGZvHK2J1Tivh_z0/edit?usp=sharing&ouid=105450876100971092600&rtpof=true&sd=true">Click here to go to the PPT</a>
<br>
Feature Specific Template link of the project: <a href="https://drive.google.com/file/d/1-atuWjaDJ355ayWXu7W0nE5bBJZxg0YE/view?usp=sharing">Feature Specific Template</a><br>
# Phase 2
<b>Objective</b>: Coding and debugging/testing<br> 
<b>Description:</b><br> 
•	Coding of GET/PUT/PATCH/POST/DELETE operations.<br>
•	Code Review<br>
•	Unit testing<br>
 •	Examples to demonstrate the working of library.<br>
	
<h2>Deliverable:</h2> Presentation of working of library using example programs.<br>

<h2>References:</h2>	
•	https://github.com/HewlettPackard/python-ilorest-library/<br>
•	https://github.com/HewlettPackard/python-ilorest-library/tree/master/examples/Redfish/<br>
•	https://hewlettpackard.github.io/ilo-rest-api-docs/ilo5/<br>
•	https://www.dmtf.org/standards/redfish<br>
 ## Requirements

 - Install golang

 - Clone project in gopath (~/go/src/github.com/sgkul2000/hpe-cty)

 - Get required packages (go mod tidy)
 ## Run program
  `go run cmd/main.go`

Currently we had uploaded the GET operation and soon we will be uploading the design document and code for other operations.


