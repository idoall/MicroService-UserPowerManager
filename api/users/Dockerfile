FROM alpine
ADD conf /conf
ADD mshk.users /mshk.users
RUN mkdir -p /logs
ENTRYPOINT [ "/mshk.users" ]
