FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.users /mshk.users
RUN mkdir -p /logs

ENTRYPOINT [ "/mshk.users" ]
