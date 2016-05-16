#!/usr/bin/env python
# -*- coding: utf-8 -*-

import urllib3

http = urllib3.PoolManager()

r = http.request('GET', 'http://asmodean.reverse.net/pages/tools_index.html')

r.data
print(asmodean.soup.find_all('table')[1])





