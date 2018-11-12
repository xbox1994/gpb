package template

import (
	"errors"
	"fmt"
	"os"
)

var logConfig = &Template{
	Mode:     Create,
	FilePath: `log4go.xml`,
	Content: `<logging>
  <filter enabled="true">
    <tag>stdout</tag>
    <type>console</type>
    <level>ACCESS</level>
    <exclude>mgox</exclude>
  </filter>
  <filter enabled="true">
    <tag>access</tag> <!-- the tag of accesslog MUST be access -->
    <type>file</type>
    <level>ACCESS</level><!-- the level of accesslog MUST be access -->
    <property name="filename">log/access.log</property>
    <property name="format">[%D %T] [%L] %M</property>
    <property name="rotate">true</property>
    <property name="maxsize">1000M</property>
    <property name="maxlines">1000K</property>
    <property name="daily">true</property>
  </filter>
  <filter enabled="true">
    <tag>file_finest</tag>
    <type>file</type>
    <!-- level is (:?FINEST|FINE|DEBUG|TRACE|INFO|WARNING|ERROR) -->
    <level>FINEST</level>
    <!--<exclude>github.com/yaosxi</exclude>-->
    <property name="filename">log/finest.log</property>
    <!--
       %T - Time (15:04:05 MST)
       %t - Time (15:04)
       %D - Date (2006/01/02)
       %d - Date (01/02/06)
       %L - Level (FNST, FINE, DEBG, TRAC, WARN, EROR, CRIT)
       %S - Source
       %M - Message
       It ignores unknown format strings (and removes them)
       Recommended: "[%D %T] [%L] (%S) %M"
    -->
    <property name="format">[%D %T] [%L] (%S) %M</property>
    <property name="rotate">true</property> <!-- true enables log rotation, otherwise append -->
    <property name="maxsize">1000M</property> <!-- \d+[KMG]? Suffixes are in terms of 2**10 -->
    <property name="maxlines">1000K</property> <!-- \d+[KMG]? Suffixes are in terms of thousands -->
    <property name="daily">true</property> <!-- Automatically rotates when a log message is written after midnight -->
  </filter>
  <filter enabled="true">
    <tag>file_info</tag>
    <type>file</type>
    <level>INFO</level>
    <property name="filename">log/info.log</property>
    <property name="format">[%D %T] [%L] (%S) %M</property>
    <property name="rotate">true</property>
    <property name="maxsize">1000M</property>
    <property name="maxlines">1000K</property>
    <property name="daily">true</property>
  </filter>
  <filter enabled="true">
    <tag>file_error</tag>
    <type>file</type>
    <level>ERROR</level>
    <property name="filename">log/error.log</property>
    <property name="format">[%D %T] [%L] (%S) %M</property>
    <property name="rotate">true</property>
    <property name="maxsize">1000M</property>
    <property name="maxlines">1000K</property>
    <property name="daily">true</property>
  </filter>
</logging>
`,
	StdOut: createLogConfig,
}

func init() {
	AvailableTemplates = append(AvailableTemplates, logConfig)
}

func createLogConfig(template *Template, args ...string) (err error) {
	if len(args) < 2 {
		err = errors.New(`params error`)
	} else {
		projectPath := args[0]
		projectName := args[1]
		absPath := fmt.Sprintf("%s/%s/", projectPath, projectName)
		if file, err1 := os.Create(absPath + template.FilePath); err1 == nil {
			_, err = file.Write([]byte(template.Content))
			file.Close()
		} else {
			err = err1
			fmt.Println(err)
		}
	}
	return
}
