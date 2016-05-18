#!/usr/bin/env python
# -*- coding: utf-8 -*-

import urllib3
from bs4 import BeautifulSoup
http = urllib3.PoolManager()

def collect_tool_page_urls(tool_idxpage_structure):
    soup = BeautifulSoup(tool_idxpage_structure, 'html.parser')
    return map(lambda e: e.attrs['href'], soup.find_all('table')[1].find_all('a'))

def map_tool_url_tool_page(tool_page_urls):
    tool_pages = map(lambda url: BeautifulSoup(http.request('GET', url).data, 'html.parser'), tool_page_urls)
    # b[0].find('a').attrs['href']
    return tool_pages

r = http.request('GET', 'http://asmodean.reverse.net/pages/tools_index.html')

tool_page_urls = collect_tool_page_urls(r.data)
a = map_tool_url_tool_page(list(tool_page_urls)[:3])
