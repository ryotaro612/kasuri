#!/usr/bin/env python
# -*- coding: utf-8 -*-

import urllib3
from bs4 import BeautifulSoup
http = urllib3.PoolManager()

r = http.request('GET', 'http://asmodean.reverse.net/pages/tools_index.html')

def collect_tool_url(tool_idxpage_structure):
    soup = BeautifulSoup(tool_idxpage_structure, 'html.parser')
    return list(map(lambda e: e.attrs['href'], soup.find_all('table')[1].find_all('a')))

print(collect_tool_url(r.data))
