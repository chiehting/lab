const puppeteer = require('puppeteer');

(async () => {
  // Create a browser instance
  const browser = await puppeteer.launch({
    headless: false,
    args: ['--no-sandbox']
  });

  // Create a new page
  const page = await browser.newPage();

  // Set viewport width and height
  await page.setViewport({ width: 1280, height: 720 });

  var website_url = 'https://google.com';
  await page.goto(website_url, { waitUntil: 'networkidle0' });
  await page.screenshot({
    path: 'google.com.png',
  });

  website_url = 'https://www.youtube.com';
  await page.goto(website_url, { waitUntil: 'networkidle0' });
  await page.screenshot({
    path: 'youtube.com.png',
  });

  website_url = 'https://medium.com';
  await page.goto(website_url, { waitUntil: 'networkidle0' });
  await page.screenshot({
    path: 'medium.com.png',
  });

  website_url = 'https://github.com';
  await page.goto(website_url, { waitUntil: 'networkidle0' });
  await page.screenshot({
    path: 'github.com.png',
  });

  // Close the browser instance
  await browser.close();
})();

