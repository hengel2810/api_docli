FROM scratch
COPY $HOME/gopath/bin/api_docli /
EXPOSE 8000
ENTRYPOINT ["/api_docli"]