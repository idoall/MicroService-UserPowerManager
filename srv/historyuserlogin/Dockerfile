FROM idoall/alpine:3.9.3
ADD conf /conf
ADD mshk.historyuserlogin /mshk.historyuserlogin
RUN mkdir -p /logs
    
ENTRYPOINT [ "/mshk.historyuserlogin" ]
