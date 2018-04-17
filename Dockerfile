FROM scratch
RUN mkdir shared
COPY ./api_docli /
EXPOSE 8000
ENTRYPOINT ["/api_docli"]