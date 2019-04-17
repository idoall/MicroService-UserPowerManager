FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.srvusers /mshk.srvusers
RUN mkdir -p /logs

ENTRYPOINT [ "/mshk.srvusers" ]
