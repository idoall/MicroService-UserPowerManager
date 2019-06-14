FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.srvhistoryuserlogin /mshk.srvhistoryuserlogin
RUN mkdir -p /logs
    
ENTRYPOINT [ "/mshk.srvhistoryuserlogin" ]
