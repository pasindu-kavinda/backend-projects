#!/bin/bash
# Author: [Pasindu Kavinda]

difficultyLevel='easy'
randomNumber=0
attemptCount=0
guessedNumber=0
maxAttempts=0

welcome() {
    gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "🎉 Hello, there! 🎉  🌟 Welcome to the  $(gum style --foreground 212 'Number Guessing Game 🌟')."
    sleep 1
    echo ""
    gum spin --spinner dot --title "🎮✨ Game Loading... Please Wait! ✨🎮" -- sleep 3
    echo ""
}

askDifficultyLevel() {
    clear
    echo -e "Choose your difficulty level:"
    difficultyLevel=$(gum choose "easy" "medium" "hard")

    case $difficultyLevel in
    easy) maxAttempts=10 ;;
    medium) maxAttempts=5 ;;
    hard) maxAttempts=3 ;;
    esac
}

generateRandomNumber() {
    randomNumber=$((RANDOM % 100 + 1))
}

askForGuess() {
    guessedNumber=$(gum input --placeholder "Guess the number")

    if [[ $guessedNumber == 'exit' ]]; then
        exitGame
    fi

    if ! [[ $guessedNumber =~ ^[0-9]+$ ]]; then
        echo "Please enter a valid number."
        echo ""
        askForGuess
    elif [[ $guessedNumber -lt 1 || $guessedNumber -gt 100 ]]; then
        echo "Please enter a number between 1 and 100."
        echo ""
        askForGuess
    fi
}

exitGame() {
    clear
    echo "👋✨ Thank You for Playing! ✨👋"
    echo "🌟 Goodbye and See You Next Time! 🌟"
    echo ""
    echo "🎮 Keep Guessing, Keep Winning! 🎯"

    gum spin --spinner dot --title "Exiting game..." -- sleep 1
    exit
}

askForRestart() {
    echo -e "Do you want to play again? (yes/no)"
    restartChoice=$(gum choose "yes" "no")

    if [[ $restartChoice == 'yes' ]]; then
        askDifficultyLevel
        startGame
    else
        exitGame
    fi
}

attempt() {
    while [[ $attemptCount -lt $maxAttempts ]]; do
        attemptCount=$((attemptCount + 1))
        gum style --foreground 212 "Attempt: $attemptCount"

        askForGuess
        gum spin --spinner dot --title "Please wait...checking your luck." -- sleep 1

        if [[ $guessedNumber -eq $randomNumber ]]; then
            gum style --border double --margin "1" --padding "1 2" --border-foreground 212 "🎉 Congratulations! You guessed the correct number in $(gum style --foreground 212 $attemptCount) attempt(s). 🎉"
            echo ""
            askForRestart
            return
        elif [[ $guessedNumber -gt $randomNumber ]]; then
            echo "Incorrect! The number is less than $guessedNumber."
        else
            echo "Incorrect! The number is greater than $guessedNumber."
        fi
        echo ""
    done

    clear
    gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "You have reached the maximum number of attempts. The correct number was $(gum style --foreground 212 $randomNumber)."
    echo ""
    askForRestart
}

startGame() {
    clear
    attemptCount=0
    gum style --border normal --margin "1" --padding "1 2" --border-foreground 212 "🎮 Starting the game with $(gum style --foreground 212 $difficultyLevel) difficulty...You have $(gum style --foreground 212 $maxAttempts) attempts to guess the correct number 🎯."
    sleep 0.5
    echo ""

    generateRandomNumber
    attempt
}

welcome
askDifficultyLevel
startGame
