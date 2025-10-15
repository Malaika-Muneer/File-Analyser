# Use a lightweight Nginx image as the base
FROM nginx:alpine

# Copy static frontend files to Nginx web directory
COPY . /usr/share/nginx/html

# Expose Nginx default port
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
