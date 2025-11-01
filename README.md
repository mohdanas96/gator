### RSS feeder built using go (Gator)
---
Stay up to date with latest content of your favourite blogs through a cli application built using go.
Gator does the tiring, boring task of keeping up with updated blog contents by scraping feeds and feeding you with the right content.
Just tell gator which blog needs to be kept track of and chill.
Multiple user regisration is supported and each user can view and follow other user's blogs as well.

## Requirements
- Go toolchain
- Postgresql

## Installation
- git clone the repository to your local computer
- Go to the root directory of the and run go install
- Create a .gatorconfig.json file in the home directory "~"
 
 ## Usage 
 - run gator agg in a another terminal session to scrape feeds continously

 ## Commands
 - gator register [username] - To register a user
 - gator login [username] - To login as a user
 - gator reset - To delete all users
 - gator users - To list current and all users
 - gator addfeed - [feed name][feed url] - To add a feed
 - gator feeds - To list all available feeds
 - gator follow [url] - To follow a specific feed
 - gator following - To list all feeds current users is following
 - gator unfollow [url] - To unfollow a feed
 - gator browse [limit] - To browse all the feeds the current user is following. 'limit' can be specified for a specific amount of posts to browse
 default is 2
 


