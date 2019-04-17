FROM alpine
ADD conf /conf
ADD mshk.apiusers /mshk.apiusers
RUN mkdir -p /logs
ENTRYPOINT [ "/mshk.apiusers" ]
