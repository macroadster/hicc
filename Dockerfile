FROM centos:7
add build/hicc /bin/hicc
add web /web
ENTRYPOINT ["/bin/hicc"]
