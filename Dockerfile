FROM alpine:3.5

ENV TIL_PORT 3000
ENV TIL_SECRET PLEASE_SET_A_SECRET
ENV TIL_REPO_URL INVALID_REPO_URL
ENV TIL_POST_DIR content/til

RUN apk add --no-cache \
    git

ADD ./til-system /usr/local

ENTRYPOINT ["/usr/local/til-system"]
CMD /usr/local/til-system


