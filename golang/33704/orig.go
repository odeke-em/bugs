package main

func GuardAthenaService(name string) {

	for true {

		cur_dir, err := oss_utils.GetCurrentDirectory()
		if err != nil {
			oss_logger.Errorf("Fail to get current runing dir, err_desc=%s", err.Error())
			continue
		}

		bin := filepath.Join(cur_dir, name)
		conf := filepath.Join(cur_dir, fmt.Sprintf("../config/%s.yml", name))
		cmd := exec.Command(bin, "-c", conf)
		oss_logger.Infof("Athena start %v", cmd)
		cmdOutput, err := cmd.CombinedOutput()
		if err != nil {
			oss_logger.Errorf("%s service out put %s", name, cmdOutput)
			oss_logger.Errorf("%s service failed:%s,service restart", name, err.Error())
		}
		time.Sleep(5 * time.Second)
	}
}

func main() {

	// ...........
	go GuardAthenaService("athena-keeper")
	go GuardAthenaService("athena-sentinel")

	// .............
}
