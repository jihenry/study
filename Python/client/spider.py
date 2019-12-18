import urllib.request
from http import cookiejar
from bs4 import BeautifulSoup
import re,time
# import html_downloader,html_parser,html_outputer
import re,sys
# sys.setrecursionlimit(10500)
class UrlManager():
	def __init__(self):
		self.new_urls = set()
		self.old_urls = set()

	def add_new_url(self,url):
		if url is None:
			return
		if url not in self.new_urls and url not in self.old_urls:
			self.new_urls.add(url)

	def add_new_urls(self,urls):
		if urls is None or len(urls) == 0:
			return
		for url in urls:
			self.add_new_url(url)

	def has_new_url(self):
		return len(self.new_urls) != 0

	def get_new_url(self):
		new_url = self.new_urls.pop()
		self.old_urls.add(new_url)
		return new_url

class HtmlParser():
	def __init__(self):
		pass

	def _get_new_urls(self,page_url, soup):
		new_urls = set()
		links = soup.find_all('a', href=re.compile(r'/touxiang/\w+/2019/\d+\.html'))
		for link in links:
			new_url = link['href']
			new_full_url = urllib.parse.urljoin(page_url,new_url)
			new_urls.add(new_full_url)
		return new_urls

	def parse(self, page_url, html_cont):
		if page_url is None or html_cont is None:
			return
		soup = BeautifulSoup(html_cont, 'html.parser')
		new_urls = self._get_new_urls(page_url, soup)
		return new_urls

class HtmlOutputer():
	def __init__(self):
		self.html = []

	def collect_data(self, html):
		print (html)
		if html is None:
			return
		self.html.append(html)

	def get_img(self, html):
		p = (r'<a href="([^<]*\.jpeg)"')
		# imglist = re.findall(p, self.html)
		for each in re.findall(p, html):
			filename = each.split("/")[-1]
			try:
				urllib.request.urlretrieve("http:"+each, "nv/"+filename, None)
			except:
				pass

class HtmlDownloader():
	def __init__(self):
		pass

	def download(self,url):
		if url is None:
			return
		req = urllib.request.Request(url)
		req.add_header('User-Agent', 'Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36')
		response = urllib.request.urlopen(req)
		if response.getcode() != 200:
			return None
		html = response.read().decode('utf-8')
		return html

class SpiderMain():

	count = 1
	def __init__(self):
		self.urls = UrlManager()
		self.downloader = HtmlDownloader()
		self.parser = HtmlParser()
		self.outputer = HtmlOutputer()
	def craw(self, root_url):
		count = 1
		self.urls.add_new_url(root_url)
		while self.urls.has_new_url():
			try:
				new_url = self.urls.get_new_url()
				print ('craw %d: %s' % (count, new_url))
				html_cont = self.downloader.download(new_url)
				new_urls = self.parser.parse(new_url, html_cont)
				self.urls.add_new_urls(new_urls)

				if count == 1500:
					break
				count = count + 1
			except Exception as e:
				# time.sleep(1)
				print ('craw failed ' + str(e))
			print ('Downloading images...')
			self.outputer.get_img(html_cont)

if __name__ == '__main__':
	root_url = 'http://www.woyaogexing.com/'
	obj_spider = SpiderMain()
	obj_spider.craw(root_url)
	print("finish.")