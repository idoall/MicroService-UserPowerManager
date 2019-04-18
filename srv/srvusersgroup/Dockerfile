FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.srvusersgroup /mshk.srvusersgroup
RUN mkdir -p /logs

ENTRYPOINT [ "/mshk.srvusersgroup" ]
