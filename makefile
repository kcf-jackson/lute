build: 
	cd /workspaces/lute/javascript/ && \
		export PATH=$$PATH:/usr/local/go/bin:$$HOME/go/bin && \
		gopherjs build --tags javascript -o lute.min.js -m && \
		cp lute.min.js /workspaces/lute/demo/