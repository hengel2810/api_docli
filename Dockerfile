FROM scratch
COPY ./api_docli /
EXPOSE 8000
ENTRYPOINT ["/api_docli"]