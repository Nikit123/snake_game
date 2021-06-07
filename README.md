# snake_game

## Assumptions and Instructions

1. This game has used only signed integers for visualization, here:

- -1 ===> Food
-  0 ===> Empty Space
-  1 ===> Snake's Head
-  any +integer after 1 ===> part of Snake's body

2. Dimensions for the 2D play area are user defined, although these cannot be less than 2 or greater than 100 for better visualizing inside console.

3. User can use usual w/a/s/d keys to move the snake around, where:
- w ===> Up
- a ===> Left
- s ===> Down
- d ===> Right

## Win and Lose

### GAME OVER

- Snake goes out of play area bounds/hits a wall.
- Snake collides with itself/tries to go back on itself.

### GAME WON

- Snake occupies all the empty spaces (by eating food and growing).
