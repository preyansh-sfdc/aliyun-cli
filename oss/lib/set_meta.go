package lib

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync/atomic"
	"time"

	oss "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var headerOptionMap = map[string]interface{}{
	oss.HTTPHeaderContentType:                  oss.ContentType,
	oss.HTTPHeaderCacheControl:                 oss.CacheControl,
	oss.HTTPHeaderContentDisposition:           oss.ContentDisposition,
	oss.HTTPHeaderContentEncoding:              oss.ContentEncoding,
	oss.HTTPHeaderExpires:                      oss.Expires,
	oss.HTTPHeaderAcceptEncoding:               oss.AcceptEncoding,
	oss.HTTPHeaderOssServerSideEncryption:      oss.ServerSideEncryption,
	oss.HTTPHeaderOssObjectACL:                 oss.ObjectACL,
	oss.HTTPHeaderOrigin:                       oss.Origin,
	oss.HTTPHeaderOssStorageClass:              oss.ObjectStorageClass,
	oss.HTTPHeaderOssServerSideEncryptionKeyID: oss.ServerSideEncryptionKeyID,
	oss.HTTPHeaderOssServerSideDataEncryption:  oss.ServerSideDataEncryption,
	oss.HTTPHeaderSSECAlgorithm:                oss.SSECAlgorithm,
	oss.HTTPHeaderSSECKey:                      oss.SSECKey,
	oss.HTTPHeaderSSECKeyMd5:                   oss.SSECKeyMd5,
}

func formatHeaderString(hopMap map[string]interface{}, sep string) string {
	str := ""
	for header := range hopMap {
		if header == oss.HTTPHeaderExpires {
			str += header + fmt.Sprintf("(time.RFC3339: %s)", time.RFC3339) + sep
		} else {
			str += header + sep
		}
	}
	if len(str) >= len(sep) {
		str = str[:len(str)-len(sep)]
	}
	return str
}

func fetchHeaderOptionMap(hopMap map[string]interface{}, name string) (interface{}, error) {
	for header, f := range hopMap {
		if strings.ToLower(name) == strings.ToLower(header) {
			return f, nil
		}
	}
	return nil, fmt.Errorf("unsupported header: %s, please check", name)
}

func getOSSOption(hopMap map[string]interface{}, name string, param string) (oss.Option, error) {
	if f, err := fetchHeaderOptionMap(hopMap, name); err == nil {
		switch f.(type) {
		case func(string) oss.Option:
			return f.(func(string) oss.Option)(param), nil
		case func(oss.ACLType) oss.Option:
			return f.(func(oss.ACLType) oss.Option)(oss.ACLType(param)), nil
		case func(t time.Time) oss.Option:
			val, err := time.Parse(http.TimeFormat, param)
			if err != nil {
				val, err = time.Parse(time.RFC3339, param)
				if err != nil {
					return nil, err
				}
			}
			return f.(func(time.Time) oss.Option)(val), nil
		case func(oss.StorageClassType) oss.Option:
			return f.(func(oss.StorageClassType) oss.Option)(oss.StorageClassType(param)), nil
		default:
			return nil, fmt.Errorf("error option type, internal error")
		}
	}
	return nil, fmt.Errorf("unsupported header: %s, please check", name)
}

var specChineseSetMeta = SpecText{

	synopsisText: "??????????????????objects????????????",

	paramText: "cloud_url [meta] [options]",

	syntaxText: ` 
    ossutil set-meta oss://bucket[/prefix] [header:value#header:value...] [--update] [--delete] [-r] [-f] [-c file] [--version-id versionId]
`,

	detailHelpText: ` 
    ????????????????????????????????????????????????objects???meta??????????????????--recursive????????????ossutil
    ?????????????????????cloud_url?????????objects?????????????????????objects???meta?????????????????????????????????
    object????????????????????????object????????????ossutil????????????

    ???1??????????????????????????????????????????--update?????????--delete?????????ossutil???????????????objects???
        meta??????????????????[header:value#header:value...]????????????[header:value#header:value...]
        ?????????????????????????????????meta??????????????????????????????headers???????????????` + oss.HTTPHeaderOssMetaPrefix + `?????????headers???
        ??????????????????????????????ossutil??????????????????????????????????????????meta?????????

    ???2?????????meta?????????????????????--update?????????ossutil???????????????objects?????????header?????????
        ???value????????????value?????????????????????objects?????????meta????????????????????????????????????--delete
        ?????????

    ???3?????????meta?????????????????????--delete?????????ossutil???????????????objects?????????header???????????????
        ?????????headers???????????????` + oss.HTTPHeaderOssMetaPrefix + `?????????headers???????????????????????????????????????value??????
        ?????????header:??????header????????????objects?????????meta????????????????????????????????????--update?????????

    ??????????????????bucket???meta?????????????????????bucket???meta??????????????????bucket???????????????
    ??????bucket??????object???meta??????????????????stat?????????

Headers:

    ?????????header???????????????
        ` + formatHeaderString(headerOptionMap, "\n        ") + `
        ?????????` + oss.HTTPHeaderOssMetaPrefix + `?????????header

    ?????????header????????????????????????value??????????????????

?????????

    ???????????????????????????

    1) ossutil set-meta oss://bucket/object [header:value#header:value...] [--update] [--delete] [-f] [--version-id versionId]
        ???????????????--recursive?????????ossutil?????????????????????object???meta??????????????????????????????
    ???cloud_url???????????????????????????meta???object??????object???????????????????????????????????????--force
    ????????????????????????????????????????????????????????????[header:value#header:value...]??????????????????
    object?????????meta???
        --update?????????--delete??????????????????????????????

    2) ossutil set-meta oss://bucket[/prefix] [header:value#header:value...] -r [--update] [--delete] [-f]
        ???????????????--recursive?????????ossutil???????????????????????????cloud_url???objects???????????????
    ??????objects???meta??????????????????object?????????????????????????????????object????????????????????????report
    ??????????????????????????????object??????????????????object???????????????????????????report????????????????????????
    ???cp?????????????????????
        ???????????????--include/--exclude?????????ossutil?????????????????????pattern???objects??????????????????
        --include???--exclude????????????????????????cp???????????????
        ??????--force????????????????????????????????????????????????
        --update?????????--delete??????????????????????????????
`,

	sampleText: ` 
    (1)ossutil set-meta oss://bucket1/obj1 Cache-Control:no-cache#Content-Encoding:gzip#X-Oss-Meta-a:b
        ??????obj1???Cache-Control???Content-Encoding???X-Oss-Meta-a??????

    (2)ossutil set-meta oss://bucket1/o X-Oss-Meta-empty:#Content-Type:plain/text --update -r
        ???????????????o?????????objects???X-Oss-Meta-empty???Content-Type??????

    (3)ossutil set-meta oss://bucket1/ X-Oss-Meta-empty:#Content-Type:plain/text --update -r --include "*.jpg"
        ?????????????????????.jpg???objects???X-Oss-Meta-empty???Content-Type??????

    (4)ossutil set-meta oss://bucket1/o X-Oss-Meta-empty:#Content-Type:plain/text --update -r --exclude "*.jpg"
        ???????????????o???????????????.jpg???objects???X-Oss-Meta-empty???Content-Type??????

    (5)ossutil set-meta oss://bucket1/obj1 X-Oss-Meta-delete --delete
        ??????obj1???X-Oss-Meta-delete??????

    (6)ossutil set-meta oss://bucket/o -r
        ???????????????o?????????objects???meta??????

    (7)ossutil set-meta oss://bucket1/%e4%b8%ad%e6%96%87 X-Oss-Meta-delete --delete --encoding-type url
        ??????oss://bucket1/?????????X-Oss-Meta-delete??????

    (6)ossutil set-meta oss://bucket1/obj1 X-Oss-Meta-delete --delete --version-id versionId
        ??????????????????obj1???X-Oss-Meta-delete??????????????????????????????
`,
}

var specEnglishSetMeta = SpecText{

	synopsisText: "set metadata on already uploaded objects",

	paramText: "cloud_url [meta] [options]",

	syntaxText: ` 
    ossutil set-meta oss://bucket[/prefix] [header:value#header:value...] [--update] [--delete] [-r] [-f] [-c file] [--version-id versionId]
`,

	detailHelpText: ` 
    The command can be used to set, update or delete the specified objects' meta data. 
    If --recursive option is specified, ossutil find all matching objects and batch set 
    meta on these objects, else, ossutil set meta on single object, if the object not 
    exist, error happens. 

    (1) Set full meta: If --update option and --delete option is not specified, ossutil 
        will set the meta of the specified objects to [header:value#header:value...], what
        user inputs. If [header:value#header:value...] is missing, it means clear the meta 
        data of the specified objects(to those headers which can not be deleted, that is, 
        the headers do not start with: ` + oss.HTTPHeaderOssMetaPrefix + `, the value will not be changed), at the 
        time ossutil will ask user to confirm the input.

    (2) Update meta: If --update option is specified, ossutil will update the specified 
        headers of objects to the values that user inputs(the values can be empty), other 
        meta data of the specified objects will not be changed. --delete option is not 
        supported in the usage. 

    (3) Delete meta: If --delete option is specified, ossutil will delete the specified 
        headers of objects that user inputs(to those headers which can not be deleted, 
        that is, the headers do not start with: ` + oss.HTTPHeaderOssMetaPrefix + `, the value will not be changed), 
        in this usage the value must be empty(like header: or header), other meta data 
        of the specified objects will not be changed. --update option is not supported 
        in the usage.

    The meta data of bucket can not be setted by the command, please use other commands. 
    User can use stat command to check the meta information of bucket or objects.

Headers:

    ossutil supports following headers:
        ` + formatHeaderString(headerOptionMap, "\n        ") + `
        and headers starts with: ` + oss.HTTPHeaderOssMetaPrefix + `

    Warning: headers are case-insensitive, but value are case-sensitive.

Usage:

    There are two usages:

    1) ossutil set-meta oss://bucket/object [header:value#header:value...] [--update] [--delete] [-f] [--version-id versionId]
        If --recursive option is not specified, ossutil set meta on the specified single 
    object. In the usage, please make sure cloud_url exactly specified the object you want to 
    set meta on, if object not exist, error occurs. If --force option is specified, ossutil 
    will not show prompt question. 
        The usage of --update option and --delete option is showed in detailHelpText. 

    2) ossutil set-meta oss://bucket[/prefix] [header:value#header:value...] -r [--update] [--delete] [-f]
        If --recursive option is specified, ossutil will search for prefix-matching objects 
    and set meta on these objects. If an error occurs, ossutil will record the error message 
    to report file, and ossutil will continue to attempt to set acl on the remaining objects(
    more information see help of cp command). 
        If --include/--exclude option is specified, ossutil will search for pattern-matching objects and 
    set meta on those objects. 
	    --include and --exclude option, please refer cp command help.
        If --force option is specified, ossutil will not show prompt question.
        The usage of --update option and --delete option is showed in detailHelpText.
`,

	sampleText: ` 
    (1)ossutil set-meta oss://bucket1/obj1 Cache-Control:no-cache#Content-Encoding:gzip#X-Oss-Meta-a:b
        Set Cache-Control, Content-Encoding and X-Oss-Meta-a header for obj1

    (2)ossutil set-meta oss://bucket1/o X-Oss-Meta-empty:#Content-Type:plain/text -u -r
        Batch update X-Oss-Meta-empty and Content-Type header on objects that start with o

    (3)ossutil set-meta oss://bucket1/ X-Oss-Meta-empty:#Content-Type:plain/text --update -r --include "*.jpg"
        Batch update X-Oss-Meta-empty and Content-Type header on objects ending with .jpg

    (4)ossutil set-meta oss://bucket1/o X-Oss-Meta-empty:#Content-Type:plain/text --update -r --exclude ".jpg"
        Batch update X-Oss-Meta-empty and Content-Type header on objects starting with o and ending with .jpg

    (5)ossutil set-meta oss://bucket1/obj1 X-Oss-Meta-delete -d
        Delete X-Oss-Meta-delete header of obj1 

    (6)ossutil set-meta oss://bucket/o -r
        Batch set the meta of objects that start with o to empty

    (7)ossutil set-meta oss://bucket1/%e4%b8%ad%e6%96%87 X-Oss-Meta-delete --delete --encoding-type url
        Delete X-Oss-Meta-delete header of oss://bucket1/??????
    
	(8)ossutil set-meta oss://bucket1/obj1 X-Oss-Meta-delete --delete --version-id versionId
        Delete X-Oss-Meta-delete header of a specific version of obj1???and generate the latest version obj1
`,
}

// SetMetaCommand is the command set meta for object
type SetMetaCommand struct {
	monitor   Monitor //Put first for atomic op on some fileds
	command   Command
	smOption  batchOptionType
	filters   []filterOptionType
	skipCount uint64
}

var setMetaCommand = SetMetaCommand{
	command: Command{
		name:        "set-meta",
		nameAlias:   []string{"setmeta", "set_meta"},
		minArgc:     1,
		maxArgc:     2,
		specChinese: specChineseSetMeta,
		specEnglish: specEnglishSetMeta,
		group:       GroupTypeNormalCommand,
		validOptionNames: []string{
			OptionRecursion,
			OptionUpdate,
			OptionDelete,
			OptionForce,
			OptionEncodingType,
			OptionInclude,
			OptionExclude,
			OptionConfigFile,
			OptionEndpoint,
			OptionAccessKeyID,
			OptionAccessKeySecret,
			OptionSTSToken,
			OptionProxyHost,
			OptionProxyUser,
			OptionProxyPwd,
			OptionRetryTimes,
			OptionRoutines,
			OptionLanguage,
			OptionOutputDir,
			OptionLogLevel,
			OptionVersionId,
			OptionPassword,
			OptionMode,
			OptionECSRoleName,
			OptionTokenTimeout,
			OptionRamRoleArn,
			OptionRoleSessionName,
			OptionReadTimeout,
			OptionConnectTimeout,
			OptionSTSRegion,
			OptionSkipVerfiyCert,
			OptionUserAgent,
		},
	},
}

// function for FormatHelper interface
func (sc *SetMetaCommand) formatHelpForWhole() string {
	return sc.command.formatHelpForWhole()
}

func (sc *SetMetaCommand) formatIndependHelp() string {
	return sc.command.formatIndependHelp()
}

// Init simulate inheritance, and polymorphism
func (sc *SetMetaCommand) Init(args []string, options OptionMapType) error {
	return sc.command.Init(args, options, sc)
}

// RunCommand simulate inheritance, and polymorphism
func (sc *SetMetaCommand) RunCommand() error {
	sc.monitor.init("Setted meta on")

	isUpdate, _ := GetBool(OptionUpdate, sc.command.options)
	isDelete, _ := GetBool(OptionDelete, sc.command.options)
	recursive, _ := GetBool(OptionRecursion, sc.command.options)
	force, _ := GetBool(OptionForce, sc.command.options)
	routines, _ := GetInt(OptionRoutines, sc.command.options)
	language, _ := GetString(OptionLanguage, sc.command.options)
	language = strings.ToLower(language)
	encodingType, _ := GetString(OptionEncodingType, sc.command.options)
	versionId, _ := GetString(OptionVersionId, sc.command.options)

	var res bool
	res, sc.filters = getFilter(os.Args)
	if !res {
		return fmt.Errorf("--include or --exclude does not support format containing dir info")
	}

	if !recursive && len(sc.filters) > 0 {
		return fmt.Errorf("--include or --exclude only work with --recursive")
	}

	if recursive && len(versionId) > 0 {
		return fmt.Errorf("--version-id only work on single object")
	}

	cloudURL, err := CloudURLFromString(sc.command.args[0], encodingType)
	if err != nil {
		return err
	}

	if err = sc.checkArgs(cloudURL, recursive, isUpdate, isDelete); err != nil {
		return err
	}

	if !sc.confirmOP(recursive, force) {
		return nil
	}

	if err := sc.checkOption(isUpdate, isDelete, force, language); err != nil {
		return err
	}

	str, err := sc.getMetaData(force, language)
	if err != nil {
		return err
	}

	headers, err := sc.command.parseHeaders(str, isDelete)
	if err != nil {
		return err
	}

	bucket, err := sc.command.ossBucket(cloudURL.bucket)
	if err != nil {
		return err
	}

	if !recursive {
		err = sc.setObjectMeta(bucket, cloudURL.object, headers, isUpdate, isDelete, versionId)
	} else {
		err = sc.batchSetObjectMeta(bucket, cloudURL, headers, isUpdate, isDelete, force, routines)
	}

	if isUpdate {
		LogInfo("update skip count:%d\n", sc.skipCount)
	}
	return err
}

func (sc *SetMetaCommand) checkArgs(cloudURL CloudURL, recursive, isUpdate, isDelete bool) error {
	if cloudURL.bucket == "" {
		return fmt.Errorf("invalid cloud url: %s, miss bucket", sc.command.args[0])
	}
	if !recursive && cloudURL.object == "" {
		return fmt.Errorf("set object meta invalid cloud url: %s, object empty. Set bucket meta is not supported, if you mean batch set meta on objects, please use --recursive", sc.command.args[0])
	}
	if isUpdate && isDelete {
		return fmt.Errorf("--update option and --delete option are not supported for %s at the same time, please check", sc.command.args[0])
	}
	return nil
}

func (sc *SetMetaCommand) checkOption(isUpdate, isDelete, force bool, language string) error {
	if !isUpdate && !isDelete && !force {
		if language == LEnglishLanguage {
			fmt.Printf("Warning: --update option means update the specified header, --delete option means delete the specified header, miss both options means update the whole meta info, continue to update the whole meta info(y or N)? ")
		} else {
			fmt.Printf("?????????--update?????????????????????header???--delete?????????????????????header??????????????????????????????object?????????meta???????????????????????????????????????meta??????(y or N)? ")
		}
		var str string
		if _, err := fmt.Scanln(&str); err != nil || (strings.ToLower(str) != "yes" && strings.ToLower(str) != "y") {
			return fmt.Errorf("operation is canceled")
		}
		fmt.Println("")
	}
	return nil
}

func (sc *SetMetaCommand) confirmOP(recursive, force bool) bool {
	if recursive && !force {
		var val string
		fmt.Printf("Do you really mean to recursivlly set meta on objects of %s(y or N)? ", sc.command.args[0])
		if _, err := fmt.Scanln(&val); err != nil || (strings.ToLower(val) != "yes" && strings.ToLower(val) != "y") {
			fmt.Println("operation is canceled.")
			return false
		}
	}
	return true
}

func (sc *SetMetaCommand) getMetaData(force bool, language string) (string, error) {
	if len(sc.command.args) > 1 {
		return strings.TrimSpace(sc.command.args[1]), nil
	}

	if force {
		return "", nil
	}

	if language == LEnglishLanguage {
		fmt.Printf("Do you really mean the empty meta(or forget to input header:value pair)? \nEnter yes(y) to continue with empty meta, enter no(n) to show supported headers, other inputs will cancel operation: ")
	} else {
		fmt.Printf("??????????????????????????????meta????????????????????????????????????header:value??????? \n??????yes(y)?????????meta?????????????????????no(n)??????????????????headers?????????????????????????????????")
	}
	var str string
	if _, err := fmt.Scanln(&str); err != nil || (strings.ToLower(str) != "yes" && strings.ToLower(str) != "y" && strings.ToLower(str) != "no" && strings.ToLower(str) != "n") {
		return "", fmt.Errorf("unknown input, operation is canceled")
	}
	if strings.ToLower(str) == "yes" || strings.ToLower(str) == "y" {
		return "", nil
	}

	if language == LEnglishLanguage {
		fmt.Printf("\nSupported headers:\n    %s\n    And the headers start with: \"%s\"\n\nPlease enter the header:value#header:value... pair you want to set: ", formatHeaderString(headerOptionMap, "\n    "), oss.HTTPHeaderOssMetaPrefix)
	} else {
		fmt.Printf("\n?????????headers:\n    %s\n    ?????????\"%s\"?????????headers\n\n????????????????????????header:value#header:value...???", formatHeaderString(headerOptionMap, "\n    "), oss.HTTPHeaderOssMetaPrefix)
	}
	if _, err := fmt.Scanln(&str); err != nil {
		return "", fmt.Errorf("meta empty, please check, operation is canceled")
	}
	return strings.TrimSpace(str), nil
}

func (cmd *Command) parseHeaders(str string, isDelete bool) (map[string]string, error) {
	if str == "" {
		return nil, nil
	}

	headers := map[string]string{}
	sli := strings.Split(str, "#")
	for _, s := range sli {
		pair := strings.SplitN(s, ":", 2)
		name := pair[0]
		value := ""
		if len(pair) > 1 {
			value = pair[1]
		}
		if isDelete && value != "" {
			return nil, fmt.Errorf("delete meta for object do no support value for header:%s, please set value:%s to empty", name, value)
		}
		if _, err := fetchHeaderOptionMap(headerOptionMap, name); err != nil && !strings.HasPrefix(strings.ToLower(name), "x-oss-") {
			return nil, fmt.Errorf("unsupported header:%s, please try \"help %s\" to see supported headers", name, cmd.name)
		}
		headers[name] = value
	}
	return headers, nil
}

func (sc *SetMetaCommand) setObjectMeta(bucket *oss.Bucket, object string, headers map[string]string, isUpdate, isDelete bool, versionId string) error {
	allheaders := headers
	isSkip := false
	if isUpdate || isDelete {
		var options []oss.Option
		if len(versionId) > 0 {
			options = append(options, oss.VersionId(versionId))
		}

		// get object meta
		props, err := sc.command.ossGetObjectStatRetry(bucket, object, options...)
		if err != nil {
			return err
		}

		// get object acl
		objectACL, err := bucket.GetObjectACL(object, options...)
		if err != nil {
			return err
		}
		props.Set(StatACL, objectACL.ACL)

		// merge
		allheaders, isSkip = sc.mergeHeader(props, headers, isUpdate, isDelete)
		if isSkip {
			atomic.AddUint64(&sc.skipCount, uint64(1))
			return nil
		}
	}

	options, err := sc.command.getOSSOptions(headerOptionMap, allheaders)
	if err != nil {
		return err
	}
	if len(versionId) > 0 {
		options = append(options, oss.VersionId(versionId))
	}
	return sc.ossSetObjectMetaRetry(bucket, object, options...)
}

func (sc *SetMetaCommand) mergeHeader(props http.Header, headers map[string]string, isUpdate, isDelete bool) (map[string]string, bool) {
	allheaders := map[string]string{}
	for name := range props {
		if _, err := fetchHeaderOptionMap(headerOptionMap, name); err == nil || strings.HasPrefix(strings.ToLower(name), strings.ToLower(oss.HTTPHeaderOssMetaPrefix)) {
			allheaders[strings.ToLower(name)] = props.Get(name)
		}
		if strings.ToLower(name) == strings.ToLower(StatACL) {
			allheaders[strings.ToLower(oss.HTTPHeaderOssObjectACL)] = props.Get(name)
		}
	}

	if isUpdate {
		equalCount := 0
		for name, val := range headers {
			objectVal, ok := allheaders[strings.ToLower(name)]
			if ok && val == objectVal {
				equalCount += 1
			}
		}

		if equalCount == len(headers) {
			// skip update
			return allheaders, true
		}

		for name, val := range headers {
			allheaders[strings.ToLower(name)] = val
		}
	}
	if isDelete {
		for name := range headers {
			delete(allheaders, strings.ToLower(name))
		}
	}
	return allheaders, false
}

func (sc *SetMetaCommand) ossSetObjectMetaRetry(bucket *oss.Bucket, object string, options ...oss.Option) error {
	retryTimes, _ := GetInt(OptionRetryTimes, sc.command.options)
	cpOptions := append(options, oss.MetadataDirective(oss.MetaReplace))
	for i := 1; ; i++ {
		_, err := bucket.CopyObject(object, object, cpOptions...)
		if err == nil {
			return err
		}
		if int64(i) >= retryTimes {
			return ObjectError{err, bucket.BucketName, object}
		}
	}
}

func (sc *SetMetaCommand) batchSetObjectMeta(bucket *oss.Bucket, cloudURL CloudURL, headers map[string]string, isUpdate, isDelete, force bool, routines int64) error {
	sc.smOption.ctnu = true
	outputDir, _ := GetString(OptionOutputDir, sc.command.options)

	// init reporter
	var err error
	if sc.smOption.reporter, err = GetReporter(sc.smOption.ctnu, outputDir, commandLine); err != nil {
		return err
	}
	defer sc.smOption.reporter.Clear()

	return sc.setObjectMetas(bucket, cloudURL, headers, isUpdate, isDelete, force, routines)
}

func (sc *SetMetaCommand) setObjectMetas(bucket *oss.Bucket, cloudURL CloudURL, headers map[string]string, isUpdate, isDelete, force bool, routines int64) error {
	// producer list objects
	// consumer set meta
	chObjects := make(chan string, ChannelBuf)
	chError := make(chan error, routines+1)
	chListError := make(chan error, 1)
	go sc.command.objectStatistic(bucket, cloudURL, &sc.monitor, sc.filters)
	go sc.command.objectProducer(bucket, cloudURL, chObjects, chListError, sc.filters)

	for i := 0; int64(i) < routines; i++ {
		go sc.setObjectMetaConsumer(bucket, headers, isUpdate, isDelete, chObjects, chError)
	}

	return sc.waitRoutinueComplete(chError, chListError, routines)
}

func (sc *SetMetaCommand) setObjectMetaConsumer(bucket *oss.Bucket, headers map[string]string, isUpdate, isDelete bool, chObjects <-chan string, chError chan<- error) {
	for object := range chObjects {
		err := sc.setObjectMetaWithReport(bucket, object, headers, isUpdate, isDelete)
		if err != nil {
			chError <- err
			if !sc.smOption.ctnu {
				return
			}
			continue
		}
	}

	chError <- nil
}

func (sc *SetMetaCommand) setObjectMetaWithReport(bucket *oss.Bucket, object string, headers map[string]string, isUpdate, isDelete bool) error {
	err := sc.setObjectMeta(bucket, object, headers, isUpdate, isDelete, "")
	sc.command.updateMonitor(err, &sc.monitor)
	msg := fmt.Sprintf("set meta on %s", CloudURLToString(bucket.BucketName, object))
	sc.command.report(msg, err, &sc.smOption)
	return err
}

func (sc *SetMetaCommand) waitRoutinueComplete(chError, chListError <-chan error, routines int64) error {
	completed := 0
	var ferr error
	for int64(completed) <= routines {
		select {
		case err := <-chListError:
			if err != nil {
				return err
			}
			completed++
		case err := <-chError:
			if err == nil {
				completed++
			} else {
				ferr = err
				if !sc.smOption.ctnu {
					fmt.Printf(sc.monitor.progressBar(true, errExit))
					return err
				}
			}
		}
	}
	return sc.formatResultPrompt(ferr)
}

func (sc *SetMetaCommand) formatResultPrompt(err error) error {
	fmt.Printf(sc.monitor.progressBar(true, normalExit))
	if err != nil && sc.smOption.ctnu {
		return nil
	}
	return err
}
