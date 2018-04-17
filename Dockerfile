FROM scratch
COPY dist/api_docli /
EXPOSE 8000
ENTRYPOINT ["/traefik"]