%global provider	github
%global provider_tld	com
%global project		percona
%global repo		grafana-dashboards
%global provider_prefix	%{provider}.%{provider_tld}/%{project}/%{repo}
%global import_path	%{provider_prefix}
%global commit		d4e98afbedf4af696ab68c2b8265a9c6652e5c76
%global shortcommit	%(c=%{commit}; echo ${c:0:7})
%define build_timestamp %(date -u +"%y%m%d%H%M")

Name:		%{project}-dashboards
Version:	1.9.0
Release:	5.%{build_timestamp}.%{shortcommit}%{?dist}
Summary:	Grafana dashboards for MySQL and MongoDB monitoring using Prometheus

License:	AGPLv3
URL:		https://%{provider_prefix}
Source0:	https://%{provider_prefix}/archive/%{commit}/%{repo}-%{shortcommit}.tar.gz
Source1:	qan-app-node_modules-1.8.0.tar.gz

BuildRequires:	nodejs
BuildArch:	noarch
Requires:	percona-grafana python python-requests
Provides:	percona-grafana-dashboards = %{version}-%{release}

%description
Grafana dashboards for MySQL and MongoDB monitoring using Prometheus.
This is a set of Grafana dashboards for database and system monitoring
using Prometheus datasource.
Dashboards are also a part of Percona Monitoring and Management project.


%prep
%setup -q -a 1 -n %{repo}-%{commit}


%build
pushd pmm-app
    ln -s ../node_modules node_modules
    npm run build
    rm node_modules
popd


%install
install -d %{buildroot}%{_datadir}/%{name}/pmm-app
cp -pa ./pmm-app/dist %{buildroot}%{_datadir}/%{name}/pmm-app
echo %{version} > %{buildroot}%{_datadir}/%{name}/VERSION


%files
%license LICENSE
%doc README.md LICENSE
%{_datadir}/%{name}


%changelog
* Wed Mar 14 2018 Mykola Marzhan <mykola.marzhan@percona.com> - 1.9.0-5
- use more new node_modules

* Tue Feb 13 2018 Mykola Marzhan <mykola.marzhan@percona.com> - 1.7.0-4
- PMM-2034 compile grafana app

* Mon Nov 13 2017 Mykola Marzhan <mykola.marzhan@percona.com> - 1.5.1-1
- PMM-1771 keep QAN Plugin in dashboards dir

* Mon Nov 13 2017 Mykola Marzhan <mykola.marzhan@percona.com> - 1.5.0-1
- PMM-1680 Include QAN Plugin into PMM

* Thu Feb  2 2017 Mykola Marzhan <mykola.marzhan@percona.com> - 1.1.0-1
- add build_timestamp to Release value

* Thu Dec 15 2016 Mykola Marzhan <mykola.marzhan@percona.com> - 1.0.7-1
- init version
