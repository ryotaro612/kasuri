#!/usr/bin/env python
# -*- coding: utf-8 -*-

import urllib3
from bs4 import BeautifulSoup
http = urllib3.PoolManager()

def collect_tool_page_urls(tool_idxpage_structure):
    soup = BeautifulSoup(tool_idxpage_structure, 'html.parser')
    return map(lambda e: e.attrs['href'], soup.find_all('table')[1].find_all('a'))

def map_tool_url_tool_page(tool_page_urls):
    all_tools = map(lambda url: BeautifulSoup(http.request('GET', url).data, 'html.parser'), tool_page_urls)
    return filter(lambda e: e.find_all('table')[1].find('s') == None,all_tools)
    # b[0].find('a').attrs['href']
def save_descriptions(description_structures):
    for tool_page in description_structures:
        with open(tool_page.title.contents[0] + ".html", 'w') as fs:
            fs.writelines(str(tool_page))

import wget
def save_tools(description_pages):
    for page in description_pages:
        wget.download(page.find_all('table')[1].find('a').attrs['href'])

r = http.request('GET', 'http://asmodean.reverse.net/pages/tools_index.html')

tool_page_urls = list(collect_tool_page_urls(r.data))
# a = list(map_tool_url_tool_page([tool_page_urls[0], tool_page_urls[-1]]))
a = list(map_tool_url_tool_page(tool_page_urls))
save_descriptions(a)
save_tools(a)
