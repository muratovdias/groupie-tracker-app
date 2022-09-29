## About groupie-tracker project
Groupie Trackers consists on receiving a given API and manipulate the data contained in it, in order to create a site, displaying the information.

- It will be given an [API](https://groupietrackers.herokuapp.com/api), that consists in four parts:

  - The first one, `artists`, containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members.

  - The second one, `locations`, consists in their last and/or upcoming concert locations.

  - The third one, `dates`, consists in their last and/or upcoming concert dates.

  - And the last one, `relation`, does the link between all the other parts, `artists`, `dates` and `locations`.

- Given all this you should build a user friendly website where you can display the bands info through several data visualizations (examples : blocks, cards, tables, list, pages, graphics, etc). It is up to you to decide how you will display it.
## Usage
To run this project, open the command line and type this commands:
```
$ git clone git@git.01.alem.school:muratovdias/groupie-tracker.git
$ cd groupie-tracker
$ go run cmd/main.go
```
## This project will help you learn about :

- Manipulation and storage of data.
- [JSON](https://www.json.org/json-en.html) files and format.
- HTML.
- Event creation and display.
- [Client-server](https://developer.mozilla.org/en-US/docs/Learn/Server-side/First_steps/Client-Server_overview).
