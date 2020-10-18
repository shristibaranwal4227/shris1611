# Appointy_meeting_scheduling_API


The task is to develop a basic version of meeting scheduling API. 
Meetings should have the following Attributes. 
1) Id
2) Title
3) Participants
4) Start Time
5) End Time
6) Creation Timestamp

Participants should have the following Attributes. 
1) Name
2) Email
3) RSVP (i.e. Yes/No/MayBe/Not Answered)

Develop an HTTP JSON API capable of the following operations,

1) Schedule a meeting
2) Should be a POST request
3) Use JSON request body
4)URL should be ‘/meetings’
5) Must return the meeting in JSON format
6) Get a meeting using id
7) Should be a GET request
8) Id should be in the url parameter
9) URL should be ‘/meeting/<id here>’
10) Must return the meeting in JSON format
11) List all meetings within a time frame
12) Should be a GET request
13) URL should be ‘/meetings?start=<start time here>&end=<end time here>’
14) Must return a an array of meetings in JSON format that are within the time range
15) List all meetings of a participant
16) Should be a GET request
17) URL should be ‘/meetings?participant=<email id>’
18) Must return a an array of meetings in JSON format that have the participant received in the email within the time range


Additional Constraints/Requirements:
1) The API should be developed using Go.
2) MongoDB should be used for storage.
3) Only packages/libraries listed here and here can be used..

