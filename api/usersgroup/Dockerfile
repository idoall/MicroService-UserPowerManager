FROM alpine
ADD conf /conf
ADD mshk.usersgroup /mshk.usersgroup
RUN mkdir -p /logs
ENTRYPOINT [ "/mshk.usersgroup" ]
