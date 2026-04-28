#include<iostream>
using namespace std;int n,i,j,a[9];string s;int main(){cin>>s;for(n=s.size();i<n;)j|=s[i++]^s[--n];cout<<!j;cin>>n;while(i<n)cin>>a[i++];sort(a,a+n);for(i=0;i<n;)cout<<a[i++];}
