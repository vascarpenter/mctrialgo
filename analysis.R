
# dbから読む


library(survival)
library(RMySQL)

# read env variables
drv<-dbDriver("MySQL")
dbhost <- Sys.getenv("MYSQL_HOST")
dbhost <- ifelse(dbhost=="","",dbhost)
dbuser <- Sys.getenv("MYSQL_USER")
dbuser <- ifelse(dbuser=="","oge",dbuser)
dbpass <- Sys.getenv("MYSQL_PASSWORD")
dbpass <- ifelse(dbpass=="","hogehogeA00",dbpass)
dbname <- Sys.getenv("MYSQL_DATABASE")
dbname <- ifelse(dbname=="","studydb",dbname)

conn<-dbConnect(drv,user=dbuser, password=dbpass, dbname=dbname, host=dbhost)
d<-dbReadTable(conn,"patients")
dbDisconnect(conn)

d$female<-as.numeric(d$female)
d$age<-as.numeric(d$age)
d$trialgroup <-as.numeric(d$trialgroup)
d$allowdate<-as.Date(d$allowdate,format="%Y-%m-%d")
d$startdate <-as.Date(d$startdate,format="%Y-%m-%d")
d$dropdate <-as.Date(d$dropdate,format="%Y-%m-%d")
d$deaddate <-as.Date(d$deaddate,format="%Y-%m-%d")
d$maccedate <-as.Date(d$maccedate,format="%Y-%m-%d")
d$finishdate <-as.Date(d$finishdate,format="%Y-%m-%d")

d$dropdays <- as.numeric(d$dropdate - d$startdate, unit="days")
d$maccedays <- as.numeric(d$maccedate - d$startdate, unit="days")
d$deaddays <- as.numeric(d$deaddate - d$startdate, unit="days")
d$finishdays <- as.numeric(d$finishdate - d$startdate, unit="days")
d$drop <- !is.na(d$dropdays)
d$macce <- !is.na(d$maccedays) 
d$dead <- !is.na(d$deaddays)
d$finish <- !is.na(d$finishdays) 

# fill days (merge dropdays,maccedays,finishdays)
for(i in 1:length(d$drop))
{
	if(d$drop[i]) {
		d$days[i]<- d$dropdays[i]
	}
	if(d$macce[i]) {
		d$days[i]<- d$maccedays[i]
	}
	if(d$finish[i]) {
		d$days[i]<- d$finishdays[i]
	}
}

source('TatsukiRcodeKMplot.r')


d.sf<-survfit(Surv(days, dead)~trialgroup,data=d)
png("./static/img/test.png",width=600,height=600, point=14)
kmplot(d.sf, col.surv=1:4,lty.surv=1:4)
dev.off()

