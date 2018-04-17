FROM iron/go:dev
RUN mkdir shared
COPY ./api_docli_app /
EXPOSE 8000
ENTRYPOINT ["./api_docli_app"]