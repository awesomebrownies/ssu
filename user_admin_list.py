#!/usr/bin/env python
#Retrieve list of users on ubuntu operating system
#Initiate a web scraper
#Isolate user list on website & scrape for info
import grp
groups = grp.getgrall()

users = {}

for group in groups:
    for user in group[3]:
        if(not(user in users.keys())):
            users[user] = {group[0]}
        else:
            users[user].add(user)

for keys,values in users.items():
    print(keys, end="   ")
    print("groups:", end=" [")
    x = 0
    for value in values:
        x+=1
        if(x == len(values)):
            print(value, end="")
        else:
            print(value, end=", ")
    print("]")
