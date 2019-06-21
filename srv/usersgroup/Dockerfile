FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.usersgroup /mshk.usersgroup
RUN mkdir -p /logs

ENTRYPOINT [ "/mshk.usersgroup" ]
