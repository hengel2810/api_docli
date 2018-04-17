FROM iron/go
WORKDIR /app
# Now just add the binary
ADD myapp /app/
ENTRYPOINT ["./myapp"]