FROM alpine
ADD conf /conf
ADD mshk.apiusersgroup /mshk.apiusersgroup
RUN mkdir -p /logs
ENTRYPOINT [ "/mshk.apiusersgroup" ]
