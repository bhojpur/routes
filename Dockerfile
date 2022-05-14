FROM moby/buildkit:v0.9.3
WORKDIR /routes
COPY routes README.md /routes/
ENV PATH=/routes:$PATH
ENTRYPOINT [ "/bhojpur/routes" ]