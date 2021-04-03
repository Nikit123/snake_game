# snake_game

## TimeLine

Day 1 (30/03/2021) - Brushed up Go, went through requirement docs and Idiomatic Go practises. Started creating basic structure of project on paper, installed required tools. Set up project on git.

Day 2 (02/04/2021) - Started building the base state for project, created required packages, finished with board display and snake,food initializations.

Day 3 (03/04/2021) - Building up from base, brainstormed various test cases, updated error handling, modified directories for better implementation, testing, added comments for better understanding, added make file.

## Assumptions and Instructions

This game has used only signed integers for visualization, here:

-1 ===> Food
 0 ===> Empty Space
 1 ===> Snake's Head
 any +integer after 1 ===> part of Snake's body

-Dimensions for the 2D play area are user defined, although these cannot be less than 2 or greater than 100 for better visualizing inside console.

-User can use usual w/a/s/d keys to move the snake around, where:
w ===> Up
a ===> Left
s ===> Down
d ===> Right

## Win and Lose

### GAME OVER

1.> Snake goes out of play area bounds/hits a wall.
2.> Snake collides with itself/tries to go back on itself.

### GAME WON

1.> Snake occupies all the empty spaces (by eating food and growing).