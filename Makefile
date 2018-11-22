NAME = i3_battery_converter

SRCS = ./i3_battery_alert.go

all: $(NAME)

$(NAME):
	go build -o $(NAME) $(SRCS)

clean:
	go clean

fclean: clean
	/bin/rm -fr $(NAME)

re: fclean all

run:
	go run $(SRCS)
