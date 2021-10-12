# student-verochka-bot

## develop

    git clone https://github.com/3crabs/student-verochka-bot student_bot
    cd student_bot
    go run main.go --token=<TELEGRAM_API_TOKEN> --key=<YANDEX_WEATHER_API_KEY>

## deploy

    git clone https://github.com/3crabs/student-verochka-bot student_bot
    cd student_bot
    go build -v -o bin/bot
    ./bin/bot --token=<TELEGRAM_API_TOKEN> --key=<YANDEX_WEATHER_API_KEY> &
