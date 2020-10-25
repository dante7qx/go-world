FROM centos:7.5.1804

ENV OPERATOR=/usr/local/bin/hello-operator \
    USER_UID=1001 \
    USER_NAME=hello-operator

# install operator binary
COPY build/_output/bin/hello-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
